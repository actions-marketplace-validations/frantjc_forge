package githubactions

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/frantjc/forge"
	"github.com/frantjc/forge/internal/tarutil"
)

type Uses struct {
	Path    string
	Version string
}

func (u *Uses) IsLocal() bool {
	return strings.HasPrefix(u.Path, "./") || filepath.IsAbs(u.Path) || len(strings.Split(u.Path, "/")) < 2
}

func (u *Uses) IsRemote() bool {
	return !u.IsLocal()
}

func (u *Uses) String() string {
	uses := u.Path
	if v := u.Version; v != "" {
		uses = uses + "@" + v
	}
	return uses
}

func (u *Uses) GetOwner() string {
	if u.IsRemote() {
		return strings.Split(u.Path, "/")[0]
	}

	return ""
}

func (u *Uses) GetRepository() string {
	if u.IsRemote() {
		return strings.Split(u.Path, "/")[1]
	}

	return ""
}

func (u *Uses) GetActionPath() string {
	if u.IsRemote() {
		elements := strings.Split(u.Path, "/")
		if len(elements) > 2 {
			return filepath.Join(elements[2:]...)
		}
	}

	return ""
}

func (u *Uses) GoString() string {
	return "&Uses{" + u.String() + "}"
}

// TODO regexp.
func Parse(uses string) (*Uses, error) {
	r := &Uses{}

	switch {
	case strings.HasPrefix(uses, "/"):
		r.Path = filepath.Clean(uses)
	case strings.HasPrefix(uses, "./"), strings.HasPrefix(uses, "../"):
		r.Path = "./" + filepath.Clean(uses)
	default:
		spl := strings.Split(uses, "@")
		if len(spl) != 2 {
			return nil, fmt.Errorf("parse uses: not a path or a versioned reference: %s", uses)
		}

		r.Path = filepath.Clean(spl[0])
		r.Version = spl[1]
	}

	return r, nil
}

func (u *Uses) MarshalJSON() ([]byte, error) {
	return []byte("\"" + u.String() + "\""), nil
}

func GetUsesMetadata(ctx context.Context, uses *Uses, dir string) (*Metadata, error) {
	_ = forge.LoggerFrom(ctx)

	if r, err := OpenDirectoryMetadata(filepath.Join(dir, uses.GetActionPath())); err == nil {
		return NewMetadataFromReader(r)
	}

	if uses.IsRemote() {
		metadata, rc, err := DownloadAction(ctx, uses)
		if err != nil {
			return nil, err
		}
		defer rc.Close()

		if err = tarutil.Extract(rc, dir); err != nil {
			return nil, err
		}

		return metadata, nil
	}

	return nil, fmt.Errorf("open local action: %s", dir)
}

func OpenUsesMetadata(uses *Uses) (io.Reader, error) {
	if uses.IsRemote() {
		return nil, fmt.Errorf("open remote action: %s", uses.Path)
	}

	return OpenDirectoryMetadata(filepath.Clean(uses.Path))
}

func OpenDirectoryMetadata(dir string) (_ io.Reader, err error) {
	for _, filename := range ActionYAMLFilenames {
		if f, err := os.Open(filepath.Join(dir, filename)); err == nil {
			return f, nil
		}
	}

	return nil, fmt.Errorf("open action in directory: %s", dir)
}

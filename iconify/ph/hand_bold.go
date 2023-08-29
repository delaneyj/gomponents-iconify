package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M188 44a32 32 0 0 0-8 1v-1a32 32 0 0 0-60.79-14A32 32 0 0 0 76 60v50.83a32 32 0 0 0-52 36.7C55.82 214.6 75.35 244 128 244a92.1 92.1 0 0 0 92-92V76a32 32 0 0 0-32-32Zm8 108a68.08 68.08 0 0 1-68 68c-35.83 0-49.71-14-82.48-83.14c-.14-.29-.29-.58-.45-.86a8 8 0 0 1 13.85-8l.21.35l18.68 30A12 12 0 0 0 100 152V60a8 8 0 0 1 16 0v60a12 12 0 0 0 24 0V44a8 8 0 0 1 16 0v76a12 12 0 0 0 24 0V76a8 8 0 0 1 16 0Z"/>`),
		g.Group(children),
	)
}
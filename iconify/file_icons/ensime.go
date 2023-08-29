package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ensime(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 278.456v116.325L224.561 512L450 395.228l-113.158-58.386l-112.28 58.044Zm224.561-162.131l112.71 58.833l-112.71 58.386l-112.71-58.386Zm0-116.325L0 116.325v117.219l224.561 117.22L450 233.543V117.219Z"/>`),
		g.Group(children),
	)
}
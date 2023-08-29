package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reapp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#27CCF5" d="M63.512.566S.23.566.23 63.85v221.486s63.283 1.265 63.283-62.332V79.667c0-8.738 7.084-15.818 15.818-15.818h46.831C187.545 63.85 188.81.566 188.81.566H63.512Z"/><path fill="#22E071" d="M192.026 114.941v143.332c0 8.738-7.084 15.818-15.818 15.818h-46.826c-61.384 0-62.653 63.283-62.653 63.283h125.297s63.283 0 63.283-63.283V52.605s-63.283-1.264-63.283 62.336Z"/>`),
		g.Group(children),
	)
}
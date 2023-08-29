package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Beakeralt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M991.998 64h-32v832q0 53-37.5 90.5t-90.5 37.5h-576q-53 0-90.5-37.5t-37.5-90.5V192l-128-160q0-13 9.5-22.5t22.5-9.5h960q13 0 22.5 9.5t9.5 22.5t-9.5 22.5t-22.5 9.5zm-96 0h-776l72 96v32h96q13 0 22.5 9.5t9.5 22.5t-9.5 22.5t-22.5 9.5h-96v64h32q13 0 22.5 9.5t9.5 22.5t-9.5 22.5t-22.5 9.5h-32v64h96q13 0 22.5 9.5t9.5 22.5t-9.5 22.5t-22.5 9.5h-96v64h32q13 0 22.5 9.5t9.5 22.5t-9.5 22.5t-22.5 9.5h-32v64h96q13 0 22.5 9.5t9.5 22.5t-9.5 22.5t-22.5 9.5h-96v64h32q13 0 22.5 9.5t9.5 22.5t-9.5 22.5t-22.5 9.5h-32q0 27 18.5 45.5t45.5 18.5h576q27 0 45.5-18.5t18.5-45.5V64z"/>`),
		g.Group(children),
	)
}
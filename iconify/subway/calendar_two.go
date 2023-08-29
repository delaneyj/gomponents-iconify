package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M149.3 0h-42.7v64h42.7V0zm256 0h-42.7v64h42.7V0zM234.7 384H192V192l-85.3 21.3v32l42.7-10.7V384h-42.7v21.3h128V384zm192-384v85.3h-85.3V0H170.7v85.3H85.3V0H0v512h512V0h-85.3zm42.6 469.3H42.7V128h426.7v341.3zm-149.3-64c53.3 0 85.3-42.7 85.3-106.7S373.3 192 320 192c-53.4 0-85.3 42.7-85.3 106.7s32 106.6 85.3 106.6zm0-181.3c21.3 0 32 32 32 74.7s-10.7 74.7-32 74.7s-32-32-32-74.7s10.7-74.7 32-74.7z"/>`),
		g.Group(children),
	)
}
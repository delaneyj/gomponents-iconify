package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NutOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 4V2m-7 8v4a7.004 7.004 0 0 0 5.277 6.787c.412.104.802.292 1.102.592L12 22l.621-.621c.3-.3.69-.488 1.102-.592a7.01 7.01 0 0 0 4.125-2.939M19 10v3.343"/><path d="M12 12c-1.349-.573-1.905-1.005-2.5-2c-.546.902-1.048 1.353-2.5 2c-1.018-.644-1.46-1.08-2-2c-1.028.71-1.69.918-3 1c1.081-1.048 1.757-2.03 2-3c.194-.776.84-1.551 1.79-2.21m11.654 5.997c.887-.457 1.28-.891 1.556-1.787c1.032.916 1.683 1.157 3 1c-1.297-1.036-1.758-2.03-2-3c-.5-2-4-4-8-4c-.74 0-1.461.068-2.15.192M2 2l20 20"/></g>`),
		g.Group(children),
	)
}
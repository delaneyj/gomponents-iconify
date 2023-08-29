package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CassetteBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M22 12c0 3.771 0 5.657-1.172 6.828C19.657 20 17.771 20 14 20h-4c-3.771 0-5.657 0-6.828-1.172C2 17.657 2 15.771 2 12c0-3.771 0-5.657 1.172-6.828C4.343 4 6.229 4 10 4h4c3.771 0 5.657 0 6.828 1.172c.654.653.943 1.528 1.07 2.828"/><path d="M13.5 13.75a2.25 2.25 0 1 1 4.5 0a2.25 2.25 0 0 1-4.5 0Zm-7.5 0a2.25 2.25 0 1 1 4.5 0a2.25 2.25 0 0 1-4.5 0Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m17.5 4.5l-.527 1.404c-.47 1.256-.706 1.884-1.22 2.24c-.514.356-1.184.356-2.525.356h-2.456c-1.34 0-2.011 0-2.525-.356c-.514-.356-.75-.984-1.22-2.24L6.5 4.5"/><path stroke-linecap="round" d="M12 4v4.5"/></g>`),
		g.Group(children),
	)
}
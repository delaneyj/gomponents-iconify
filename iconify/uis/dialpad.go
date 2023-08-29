package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dialpad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 9.2h-4c-.4 0-.8.3-.8.8v4c0 .4.3.8.8.8h4c.4 0 .8-.3.8-.8v-4c0-.4-.4-.8-.8-.8zm0 7h-4c-.4 0-.8.3-.8.8v4c0 .4.3.8.8.8h4c.4 0 .8-.3.8-.8v-4c0-.4-.4-.8-.8-.8zm-7-14H3c-.4 0-.8.4-.8.8v4c0 .4.3.8.8.8h4c.4 0 .8-.3.8-.8V3c0-.4-.4-.8-.8-.8zm0 7H3c-.4 0-.8.3-.8.8v4c0 .4.3.8.8.8h4c.4 0 .8-.3.8-.8v-4c0-.4-.4-.8-.8-.8zm14-7h-4c-.4 0-.8.3-.8.8v4c0 .4.3.8.8.8h4c.4 0 .8-.3.8-.8V3c0-.4-.4-.8-.8-.8zm-7 0h-4c-.4 0-.8.4-.8.8v4c0 .4.3.8.8.8h4c.4 0 .8-.3.8-.8V3c0-.4-.4-.8-.8-.8zm7 7h-4c-.4 0-.8.3-.8.8v4c0 .4.3.8.8.8h4c.4 0 .8-.3.8-.8v-4c0-.4-.4-.8-.8-.8z"/>`),
		g.Group(children),
	)
}
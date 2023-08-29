package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashlightOnBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12.75 2a.75.75 0 0 0-1.5 0v3a.75.75 0 0 0 1.5 0V2Zm-6.22.47a.75.75 0 0 0-1.06 1.06l2 2a.75.75 0 0 0 1.06-1.06l-2-2Zm12 0a.75.75 0 0 0-1.06 0l-2 2a.75.75 0 0 0 1.06 1.06l2-2a.75.75 0 0 0 0-1.06ZM20 10v.25H4V10c0-.943 0-1.414.293-1.707C4.586 8 5.057 8 6 8h12c.943 0 1.414 0 1.707.293C20 8.586 20 9.057 20 10Zm-3.828 4.828c-.163.163-.302.302-.422.427v-.005h-7.5v.005c-.12-.125-.26-.264-.421-.427L4.75 11.75h14.5l-3.078 3.078Z"/><path fill-rule="evenodd" d="M9 20v-2.343c0-.378 0-.668-.015-.907h6.03c-.015.24-.015.53-.015.907V20c0 .943 0 1.414-.293 1.707C14.414 22 13.943 22 13 22h-2c-.943 0-1.414 0-1.707-.293C9 21.414 9 20.943 9 20Zm3-1.75a.75.75 0 0 1 .75.75v1a.75.75 0 0 1-1.5 0v-1a.75.75 0 0 1 .75-.75Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
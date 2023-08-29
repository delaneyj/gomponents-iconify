package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackSy0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackSy1)"><use href="#flagpackSy0"/><path fill="#F7FCFF" fill-rule="evenodd" d="M0 0h32v24H0V0Z" clip-rule="evenodd"/><path fill="#409100" fill-rule="evenodd" d="M9.001 13.247L7.754 14l.285-1.47L7 11.432l1.406-.06L9.001 10l.595 1.372H11L9.964 12.53l.312 1.47L9 13.247Zm14 0L21.754 14l.284-1.47L21 11.432l1.406-.06l.595-1.371l.595 1.372H25l-1.036 1.159l.312 1.469L23 13.247Z" clip-rule="evenodd"/><path fill="#E31D1C" fill-rule="evenodd" d="M0 0v8h32V0H0Z" clip-rule="evenodd"/><path fill="#272727" fill-rule="evenodd" d="M0 16v8h32v-8H0Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackSy1"><use href="#flagpackSy0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}
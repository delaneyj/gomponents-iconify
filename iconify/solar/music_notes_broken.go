package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicNotesBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-width="1.5" d="M12 19.5a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Zm10-2a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="m22 8l-10 4"/><path fill="currentColor" d="m14.456 5.158l.29.692l-.29-.692Zm2-.837l-.29-.692l.29.692Zm4.652-.98l-.416.624l.416-.624ZM21.25 12a.75.75 0 0 0 1.5 0h-1.5Zm-8.5 7V8.847h-1.5V19h1.5Zm1.995-13.15l2-.837l-.579-1.384l-2 .837l.58 1.384Zm8.005 2.16c0-1.333.002-2.42-.12-3.24c-.123-.837-.4-1.583-1.106-2.054l-.832 1.249c.185.123.355.353.455 1.024c.101.686.103 1.638.103 3.022h1.5Zm-6.005-2.997c1.276-.534 2.156-.9 2.828-1.072c.657-.167.935-.099 1.12.024l.83-1.249c-.707-.47-1.502-.437-2.32-.228c-.805.205-1.806.626-3.037 1.141l.58 1.384ZM12.75 8.848c0-.662.001-1.098.037-1.434c.035-.317.095-.474.172-.59l-1.248-.83c-.258.387-.366.805-.415 1.258c-.047.436-.046.967-.046 1.596h1.5Zm1.416-4.382c-.58.243-1.07.447-1.454.659c-.4.22-.743.48-1.001.868l1.248.831c.077-.115.199-.232.478-.386c.296-.163.698-.333 1.308-.588l-.579-1.384ZM22.75 12V8.01h-1.5V12h1.5Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M7 11V2"/><circle cx="4.5" cy="10.5" r="2.5" stroke="currentColor" stroke-width="1.5"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M10 5c-1.243 0-3-.929-3-3"/></g>`),
		g.Group(children),
	)
}
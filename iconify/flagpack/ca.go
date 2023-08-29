package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ca(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackCa0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackCa1)"><use href="#flagpackCa0"/><path fill="#F7FCFF" fill-rule="evenodd" d="M8 0h16v24H8V0Z" clip-rule="evenodd"/><path fill="#E31D1C" fill-rule="evenodd" d="M15.976 7L16 19h-.6l.368-2.97c-2.69.484-3.924.604-3.701.36c.222-.244.397-.59.525-1.038L9 12.955c.378-.004.72-.098 1.028-.281c.307-.183.153-.816-.461-1.9l1.813.264l.687-.746l1.368 1.432h.615l-.615-3.28l1.103.615L15.976 7Zm0 0l1.486 2.06l1.103-.617l-.615 3.281h.615l1.368-1.432l.686.746l1.814-.264c-.614 1.084-.768 1.717-.46 1.9c.307.183.65.277 1.027.281l-3.593 2.397c.129.448.304.794.526 1.038c.223.244-1.011.124-3.701-.36l.367 2.97H16l-.024-12ZM24 0h8v24h-8V0ZM0 0h8v24H0V0Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackCa1"><use href="#flagpackCa0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}
package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ba(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackBa0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackBa1)"><use href="#flagpackBa0"/><path fill="#2E42A5" fill-rule="evenodd" d="M0 0h32v24H0V0Z" clip-rule="evenodd"/><path fill="#FECA00" fill-rule="evenodd" d="M9 0h20v24L9 0Z" clip-rule="evenodd"/><path fill="#F7FCFF" fill-rule="evenodd" d="m6.108 2.794l-1.176.618l.225-1.31l-.951-.926L5.52.985l.588-1.191l.588 1.19l1.314.192l-.95.927l.224 1.309l-1.176-.618Zm3.487 4.28l-1.176.618l.225-1.31l-.951-.926l1.314-.191l.588-1.191l.587 1.19l1.315.192l-.951.927l.224 1.309l-1.175-.618Zm3.268 4.164l-1.175.618l.224-1.309l-.95-.927l1.314-.19l.587-1.192l.588 1.191l1.314.191l-.95.927l.224 1.31l-1.176-.619Zm3.202 3.992l-1.175.617l.224-1.309l-.95-.927l1.313-.19l.588-1.192l.588 1.191l1.314.191l-.95.927l.224 1.31l-1.176-.619Zm3.449 4.107l-1.175.618l.224-1.31l-.95-.926l1.314-.191l.587-1.191l.588 1.19l1.314.192l-.95.927l.224 1.309l-1.176-.618Zm3.813 3.974l-1.176.618l.225-1.309l-.951-.927l1.314-.191l.588-1.191l.588 1.191l1.314.191l-.951.927l.225 1.31l-1.176-.619Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackBa1"><use href="#flagpackBa0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}
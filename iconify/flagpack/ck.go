package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackCk0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackCk1)"><use href="#flagpackCk0"/><path fill="#2E42A5" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><g clip-path="url(#flagpackCk2)"><path fill="#2E42A5" d="M0 0h18v14H0z"/><path fill="#F7FCFF" d="m-2.004 13l3.96 1.737L18.09 1.889l2.09-2.582l-4.236-.58l-6.58 5.536l-5.297 3.73L-2.004 13Z"/><path fill="#F50100" d="m-1.462 14.217l2.018 1.008L19.429-.933h-2.833l-18.058 15.15Z"/><path fill="#F7FCFF" d="m20.004 13l-3.96 1.737L-.09 1.889L-2.18-.693l4.236-.58l6.58 5.536l5.297 3.73L20.004 13Z"/><path fill="#F50100" d="m19.87 13.873l-2.019 1.009l-8.036-6.918l-2.383-.773L-2.38-.684H.453l9.807 7.688l2.605.927l7.004 5.942Z"/><path fill="#F50100" fill-rule="evenodd" d="M9.985 0h-2v6H0v2h7.985v6h2V8H18V6H9.985V0Z" clip-rule="evenodd"/><path fill="#F7FCFF" fill-rule="evenodd" d="M6.485 0h5v4.5H18v5h-6.515L11.5 14h-5l-.015-4.5H0v-5h6.485V0Zm1.5 6H0v2h7.985v6h2V8H18V6H9.985V0h-2v6Z" clip-rule="evenodd"/></g><path fill="#fff" fill-rule="evenodd" d="m23.667 13.855l.726-.509l.702.51l-.218-.906l.564-.574h-.718l-.331-.74l-.282.74h-.843l.645.574l-.245.905Zm2.813 1.258l.726-.508l.702.508l-.218-.905l.564-.574h-.718l-.331-.74l-.283.74h-.842l.645.574l-.245.905Zm1.565 1.872l-.726.51l.245-.906l-.645-.574h.842l.283-.74l.33.74h.719l-.564.574l.218.905l-.702-.509Zm-1.723 2.784l.726-.509l.702.51l-.218-.906l.564-.574h-.718l-.331-.74l-.282.74h-.843l.645.574l-.245.905Zm-1.822.737l-.726.51l.245-.906l-.645-.574h.842l.283-.74l.33.74h.719l-.564.574l.218.905l-.702-.509Zm-3.1-.697l.726-.51l.702.51l-.218-.906l.564-.574h-.718l-.331-.739l-.282.74H21l.645.573l-.245.906Zm-.54-2.807l-.726.509l.245-.906l-.645-.574h.842l.283-.739l.33.74h.719l-.564.573l.218.906l-.702-.51Zm.187-1.842l.726-.51l.702.51l-.218-.906l.564-.574h-.718l-.331-.739l-.283.74h-.842l.645.573l-.245.906Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackCk1"><use href="#flagpackCk0"/></clipPath><clipPath id="flagpackCk2"><path fill="#fff" d="M0 0h18v14H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}
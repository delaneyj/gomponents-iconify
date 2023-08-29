package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g opacity=".2"><path d="M12.31 5.353a4 4 0 1 1 6.706 4.362l-2.18 3.353a4 4 0 0 1-6.707-4.362l2.18-3.353Z"/><path fill-rule="evenodd" d="M12.392 12.564a2 2 0 0 1-.587-2.767l2.181-3.353a2 2 0 1 1 3.353 2.18l-2.18 3.354a2 2 0 0 1-2.767.586Zm-2.442.202a3.999 3.999 0 0 1 .179-4.06l2.18-3.353a4 4 0 1 1 6.707 4.362l-2.18 3.353a4 4 0 0 1-6.886-.302Z" clip-rule="evenodd"/><path d="M7.362 9.293a4 4 0 0 1 6.706 4.361l-2.18 3.353a4 4 0 0 1-6.707-4.361l2.18-3.353Z"/><path fill-rule="evenodd" d="M7.444 16.503a2 2 0 0 1-.586-2.767l2.18-3.353a2 2 0 1 1 3.354 2.18l-2.181 3.354a2 2 0 0 1-2.767.586Zm-2.442.202a3.999 3.999 0 0 1 .179-4.06l2.18-3.352a4 4 0 0 1 6.707 4.361l-2.18 3.353a4 4 0 0 1-6.886-.302Z" clip-rule="evenodd"/></g><path fill-rule="evenodd" d="M10.346 11.902a3 3 0 0 1-.879-4.15l2.18-3.354a3 3 0 1 1 5.03 3.272l-2.18 3.353a3 3 0 0 1-4.15.88Zm-1.31.193a4.002 4.002 0 0 1-.407-4.889l2.18-3.353a4 4 0 1 1 6.707 4.362l-2.18 3.353a4 4 0 0 1-6.3.527Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M5.398 15.841a3 3 0 0 1-.879-4.15L6.7 8.338a3 3 0 1 1 5.03 3.271l-2.18 3.353a3 3 0 0 1-4.15.88Zm-1.31.194a4.002 4.002 0 0 1-.407-4.89l2.18-3.352a4 4 0 1 1 6.707 4.361l-2.18 3.353a4 4 0 0 1-6.3.528Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blockinger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.7 6.91a2.2 2.2 0 0 0-2.2 2.2v2.65h4.8V6.91Zm2.6 0h4.9v4.85H9.3z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.2 6.91h4.9v4.85h-4.9zm4.9 0H24v4.85h-4.9zm4.9 0h4.9v4.85H24zm4.9 0h4.9v4.85h-4.9z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.79 6.91h4.9v4.85h-4.9zm4.9 0v4.85h4.81V9.11a2.2 2.2 0 0 0-2.2-2.2ZM4.5 11.76h4.8v4.9H4.5zm4.8 0h4.9v4.9H9.3z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.2 11.76h4.9v4.9h-4.9zm14.7 0h4.9v4.9h-4.9zm4.9 0h4.9v4.9h-4.9zm4.9 0h4.8v4.9h-4.8zM4.5 16.65h4.8v4.9H4.5z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.3 16.65h4.9v4.9H9.3z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.2 16.65h4.9v4.9h-4.9zm14.7 0h4.9v4.9h-4.9zm4.9 0h4.9v4.9h-4.9zm4.9 0h4.8v4.9h-4.8zm-34.2 4.9h4.8v4.9H4.5zm4.8 0h4.9v4.9H9.3zm4.9 0h4.9v4.9h-4.9zm4.9 0H24v4.9h-4.9zm4.9 0h4.9v4.9H24zm4.9 0h4.9v4.9h-4.9zm4.9 0h4.9v4.9h-4.9zm4.9 0h4.8v4.9h-4.8zm-34.2 4.9h4.8v4.9H4.5zm14.6 0H24v4.9h-4.9zm4.9 0h4.9v4.9H24zm14.7 0h4.8v4.9h-4.8zm-34.2 4.9h4.8v4.9H4.5zm4.81 0h4.9v4.9h-4.9zm24.49 0h4.9v4.9h-4.9zm4.9 0h4.8v4.9h-4.8z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 36.24v2.65a2.2 2.2 0 0 0 2.2 2.2h2.61v-4.84Zm4.81 0h4.9v4.85h-4.9zm4.9 0h4.9v4.85h-4.9z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.1 36.24H24v4.85h-4.9zm4.9 0h4.9v4.85H24zm4.9 0h4.9v4.85h-4.9zm4.9 0h4.9v4.85h-4.9zm4.9 0v4.85h2.6a2.2 2.2 0 0 0 2.2-2.2v-2.65Z"/>`),
		g.Group(children),
	)
}
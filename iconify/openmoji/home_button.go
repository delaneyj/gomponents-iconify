package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#D0CFCE" d="M17.129 59.738L16.061 34.74l.02-6.944L36.149 8.11l19.832 19.81l-.004 15.438l-.94 8.661l-.115 7.719l-13.136-.576l.029-17.535h-11.59l-.076 17.535z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="M41.99 59.95h11.992c.55 0 1-.45 1-1V34.014m-37.924-.001V58.95c0 .55.45 1 1 1h12.135"/><path d="M8.492 35.595L36.016 7.977l27.58 27.37M41.815 59.933V41.627h-11.59v18.306"/></g>`),
		g.Group(children),
	)
}
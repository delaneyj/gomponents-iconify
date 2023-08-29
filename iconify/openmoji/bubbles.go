package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bubbles(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round"><path stroke-width="4" d="M40 47.091A5.348 5.348 0 0 1 42.091 45"/><path stroke-width="3" d="M48.5 25.5a2.562 2.562 0 0 1 1-1"/><path stroke-width="4" d="M18 21.346A8.559 8.559 0 0 1 21.346 18"/></g><path fill="#92D3F5" d="M24.5 39C31.956 39 38 32.956 38 25.5S31.956 12 24.5 12S11 18.044 11 25.5S17.044 39 24.5 39Zm26.799-3.3a7.893 7.893 0 1 0 0-15.786a7.893 7.893 0 0 0 0 15.786ZM44 60c5.523 0 10-4.477 10-10s-4.477-10-10-10s-10 4.477-10 10s4.477 10 10 10Z"/><path fill="none" stroke="#61B2E4" stroke-linecap="round" stroke-linejoin="round" stroke-width="10" d="M32 25.5a6.5 6.5 0 0 1-6.5 6.5"/><path fill="none" stroke="#61B2E4" stroke-linecap="round" stroke-linejoin="round" stroke-width="6" d="M50 50a5 5 0 0 1-5 5"/><path fill="none" stroke="#61B2E4" stroke-linecap="round" stroke-linejoin="round" stroke-width="5" d="M55.5 28.5a3 3 0 0 1-3 3"/><path fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M40 47.091A5.348 5.348 0 0 1 42.091 45"/><path fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M48.5 25.5a2.562 2.562 0 0 1 1-1"/><path fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M18 21.346A8.559 8.559 0 0 1 21.346 18"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M25.5 39C32.956 39 39 32.956 39 25.5S32.956 12 25.5 12S12 18.044 12 25.5S18.044 39 25.5 39ZM52 36a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm-7 24c5.523 0 10-4.477 10-10s-4.477-10-10-10s-10 4.477-10 10s4.477 10 10 10Z"/><path d="M45 60c5.523 0 10-4.477 10-10s-4.477-10-10-10s-10 4.477-10 10s4.477 10 10 10Z"/></g>`),
		g.Group(children),
	)
}
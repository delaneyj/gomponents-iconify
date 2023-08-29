package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AltosOdyssey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.73 5.5c4.578 0 5.05 1.81 5.148 5.091c.097 3.28.585 8.023 0 11.66c-.585 3.639-2.209 6.075-4.45 6.075s-3.703-6.253-7.178-6.253c-.563 0-1.16.221-1.767.592c.987-2.232 2.936-4.18 2.936-4.18c-.097-2.73-1.012-12.985 5.31-12.985Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.42 18.484c1.948.812 2.484 2.783 4.53 2.783s3.768-2.068 3.865-3.822c1 .08 1.446-1.933 0-2.03c.065-.942.69-3.63-1.258-4.093c-.926 1.267-5.521.976-6.958.367"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.22 17.299c.837.17 2.396-.106 2.973-.447"/><circle cx="19.222" cy="13.888" r=".75" fill="currentColor"/><circle cx="23.266" cy="13.888" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.483 22.665c-1.96 3.2-4.007 5.563-4.348 19.52m7.285-15.808c-1.682 3.118-2.437 5.866-2.486 15.807"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.761 42.184c.147-5.259-.206-14.005-.206-14.005M27.61 23.56c.09 5.131 1.965 14.997 2.623 18.749M27.926 21.93c2.136 3.594 2.295 8.613 2.295 8.613"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m29.673 39.17l7.419.701a4.335 4.335 0 0 0 4.73-3.986h0a4.335 4.335 0 0 0-3.916-4.645l-9.66-.908m-16.267-2.814l-1.128-.055a4.468 4.468 0 0 0-4.669 4.077h0a4.468 4.468 0 0 0 4.025 4.834l.21.02m14.14-25.072c1.34-1.072 1.973-3.922 1.242-5.359M16.79 37.582c3.528.317 8.94.365 12.604 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.316 42.184c-.025-1.775-.512-4.346-.512-4.346m11.421 1.668c.979 1.78.395 3.963-3.04 2.53m-9.677-26.64c.07.189.475.21.587.098"/>`),
		g.Group(children),
	)
}
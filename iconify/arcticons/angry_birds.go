package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngryBirds(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.3 41.7c.4-1.2 3.9-5.1 7.4-6.5c2-.8 5.1-1.7 6.2-1.4"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.9 9.7C21.1 8.8 18 6.9 18 5.1c0-.8.6-1.4 1.1-1.9c2-1.3 5.3.7 6.6 1.5c1.1.7 3.6 2.9 5.3 5.3c1.8.6 6.6 2.3 10.2 7.3c1.1 1.5 4 5.7 3.6 11.6c-.1 1-.3 4.3-2.4 7.6c-4.9 7.3-15.8 8.9-23.7 6.5c-2.3-.7-6-1.8-8.5-5.3c-4.5-6.2-1.3-14.5-.6-16.8c.9-2.2 2.1-5.6 5.6-8c2.2-1.5 4.5-2.1 5.9-2.2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.3 7.2c-.8 0-6.3.1-7.3 2.5c-.1.3-.2.6-.1 1c.6 1.4 4.8 1 5.6.9M7.148 28.569L1.6 28.3l.5-2.7l5.873 1.987m0 0L4 24.7l1.6-2l2.6 3.4"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.843 29.474L5.1 31.9L4 30.1l3.916-1.905M22.8 29.9c.7 1.4-.1 3-1.1 3.2c-1 .2-2.2-.8-2.4-2.3c-.1-1.5.2-3 1.1-3.2c1.3-.2 1.8.9 2.4 2.3Z"/><ellipse cx="16.4" cy="33" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.4" ry="1.8"/><circle cx="31.2" cy="26.6" r=".75" fill="currentColor"/><circle cx="36.377" cy="26.05" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.1 20.5L34 25.3l-10.7-2l.6-2.7l9.9 3.5l8.1-5.7l1.2 2.1zm-16.6 3.6c-.1.2-1.5 2.7.3 4.7c1.1 1.2 6.3 2.2 7.1-3.2m6-2.9c1.9 2.2 1.1 4.5-.2 5.4c-1.4 1-4.354.469-5.8-2.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.2 27.7c-.9.6-1.9 1.4-2.9 2.4s-1.8 2-2.3 2.9c.6.4 1.8.6 2.5.6c4.1-.1 8.1.2 10.9-.3c-.8-1-2.1-2.7-4.3-4c-1.4-1-2.7-1.4-3.9-1.6Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.9 33.7c.2 1 .4 1.9 1 3.1c.4 1.2 1 2.2 1.4 3.1c-.7.2-1.4.4-2.2.6c-2.4.4-4.3.2-6.1 0c-.7-1.1-1.5-2.7-2-4.9c-.2-1-.2-1.2 0-2.7"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.1 33c.1.7.3 1.7 1 2.7c.2.3 1 1.4 3.6 2.7c1.2.6 3 1.2 5.5 1.5"/>`),
		g.Group(children),
	)
}
package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecycleSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M20.8 3.1c-3-1.6-6.7-.6-8.4 2.4l-2.2 3.8l-2-1.1C8.2 8 8 8 7.9 8c-.5 0-.9.4-.9.9v7.2c0 .3.1.6.4.8c.1.1.3.1.4.1c.2 0 .3 0 .4-.1l6.3-3.6c.3-.2.4-.4.4-.8c0-.3-.2-.6-.4-.8L12 10.3l2.2-3.8c.4-.7 1-1.3 1.7-1.7c2-1.1 4.5-.3 5.6 1.7c.3.5.9.6 1.4.4c.5-.3.6-.9.4-1.4c-.7-1-1.5-1.9-2.5-2.4z" class="clr-i-solid clr-i-solid-path-1"/><path fill="currentColor" d="m32.2 21.1l-3-5.3l2.3-1.3c.3-.2.4-.4.4-.8c0-.3-.2-.6-.4-.8l-6.2-3.6c-.1-.1-.3-.1-.4-.1c-.5 0-.9.4-.9.9v7.2c0 .3.2.6.4.8c.1.1.3.1.4.1c.2 0 .3-.1.4-.1l2.2-1.3l3 5.3c.7 1.2.7 2.8 0 4c-.7 1.2-1.9 1.9-3.2 1.9h-.9c-.6 0-1 .4-1 1s.4 1 1 1h.9c2.1 0 4-1.1 5-3c1-1.7 1-4 0-5.9z" class="clr-i-solid clr-i-solid-path-2"/><path fill="currentColor" d="m21.7 28.4l-6.2-3.6c-.1-.1-.3-.1-.4-.1c-.5 0-.9.4-.9.9v2.6H9.3c-1.5 0-2.9-.8-3.6-2.1c-.8-1.4-.8-3.1 0-4.5c.3-.5.1-1.1-.4-1.4c-.5-.3-1.1-.1-1.4.4c-1.2 2-1.2 4.5 0 6.5C5 29 7 30.2 9.3 30.2h4.8v2.6c0 .3.2.6.4.8c.1.1.3.1.4.1c.1 0 .3 0 .4-.1l6.3-3.6c.3-.2.4-.4.4-.8c.1-.4-.1-.7-.3-.8z" class="clr-i-solid clr-i-solid-path-3"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}
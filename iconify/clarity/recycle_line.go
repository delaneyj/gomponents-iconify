package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecycleLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M6.4 17.4c.2.1.3.1.5.1s.4-.1.5-.1l7-4.1c.3-.2.5-.5.5-.9s-.2-.7-.5-.9L11.9 10L14 6.2c.4-.7 1-1.3 1.7-1.7c2-1.1 4.5-.3 5.6 1.7c.3.5.9.6 1.4.3h.1c.4-.3.5-.9.3-1.3c-.6-1-1.4-1.9-2.4-2.4c-3-1.6-6.7-.6-8.3 2.4L9.6 9.9c-.3.5-.1 1.1.3 1.4l2 1.2l-4 2.4V8.2c0-.6-.4-1-1-1c-.5.1-.9.5-.9 1.1v8.3c0 .3.2.6.4.8z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="m32.1 21l-3.5-6.2c-.1-.2-.4-.4-.6-.5c-.3-.1-.5 0-.8.1L25 15.7V11l5.5 3.3c.1 0 .1.1.2.1c.5.2 1.1 0 1.3-.5c.2-.5 0-1.1-.5-1.3l-7-4.2c-.3-.2-.7-.2-1 0c-.4.1-.5.4-.5.8v8.3c0 .4.1.8.4 1c.3.2.7.2 1 0l2.9-1.7l3 5.3c.7 1.3.7 2.8 0 4.1c-.6 1.2-1.9 1.9-3.2 1.9h-.9c-.5 0-1.2.4-1.2 1c.1.6.6 1 1.2 1h.9c2.1 0 4-1.1 5-2.9c1.1-2 1.1-4.3 0-6.2z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="m22.4 28.2l-7-4.2c-.3-.2-.7-.2-1 0c-.3.2-.4.5-.4.9v3.3H9.1c-1.5-.1-2.9-.9-3.6-2.3c-.8-1.4-.8-3.2 0-4.6c.3-.5.1-1.1-.4-1.4c-.5-.3-1.1-.1-1.4.4c-1.2 2.1-1.1 4.6.1 6.6C4.9 28.8 7 30 9.2 30H15c.6 0 1-.4 1-1v-2.4l4 2.4l-5.6 3.3c-.3.2-.5.5-.5.9c0 .6.5 1 1 1c.2 0 .3-.1.5-.2l7-4.2c.2-.1.3-.2.4-.4c.3-.4.1-1-.4-1.2z" class="clr-i-outline clr-i-outline-path-3"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}
package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IceHockey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#3e4347" d="M23.3 36.9c-5.5-.4-12.1-1.7-18.5-3.2c-1.9-.4-4.1 9.4-2 10.1c12.8 4 26.3 3 31.3 2.2c1.8-.3 3.1-2.4 3.2-2.4l8-13.4l-4.6-2.9c-7.2 8.5-13.8 9.9-17.4 9.6"/><path fill="#ed4c5c" d="M56.5 2S42.4 25.3 40.6 27.3l4.6 2.9L62 2h-5.5z"/><g fill="#3e4347"><ellipse cx="48.7" cy="57" rx="13.3" ry="5"/><ellipse cx="48.7" cy="51.1" rx="13.3" ry="5"/><path d="M35.5 51.1H62V57H35.5z"/></g><ellipse cx="48.7" cy="51.1" fill="#b2c1c0" opacity=".5" rx="13.3" ry="5"/><path fill="#e8e8e8" d="M9.1 34.6L7.5 45.5c8.5 1.9 16.7 1.7 22.2 1.2l1.8-12c-3.3 1.8-6.2 2.2-8.1 2.1c-4.3-.3-9.3-1.1-14.3-2.2"/><path fill="#b2c1c0" d="m9.2 45.8l1.5.3l2.8-10.6l-1.5-.3l-2.8 10.6m15 1.2h1.6l3.6-11.3l-1.8.6L24.2 47m-8.1-.3c.5 0 1 .1 1.5.1l.8-10.6c-.5-.1-1-.1-1.5-.2l-.8 10.7m6-10L19.9 47h1.5l2.3-10.2h-.3c-.4-.1-.8-.1-1.3-.1"/>`),
		g.Group(children),
	)
}
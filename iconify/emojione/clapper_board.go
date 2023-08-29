package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClapperBoard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#686b6d" d="m16.1 31.2l-10.9-.7V17z"/><path fill="#333" d="m4.9 11.7l6 9.8l5.4-.6l-6.4-10.1zM3.2 22.3l2.7-.3l-3.4-5.7zM15.3 9.8l7 10.4l6.5-.7l-7.6-10.8zm12.4-2.3L36 18.7l7.9-.8l-9-11.7zm31.6 8.7l-.6-4.4l-7.1-8.7l-8.8 1.6l9.9 12.2z"/><path fill="#fff" d="m21.2 8.7l7.6 10.8l7.2-.8l-8.3-11.2zM9.9 10.8l6.4 10.1l6-.7l-7-10.4zm32.9-6.1l-7.9 1.5l9 11.7l8.8-1zM2 12.2l.5 4.1L5.9 22l5-.5l-6-9.8zm56.3-4L57.6 2l-6 1.1l7.1 8.7z"/><path fill="#ebf2f2" d="m58.7 11.8l2.9 4.1l-1.7-13.7l-2.3-.2z"/><path fill="#686b6d" d="m59.3 16.2l-.6-4.4l2.9 4.1z"/><path fill="#333" d="M3 22.8v8.9l56.7 3.7V22.8z"/><path fill="#fff" d="m22.8 33l6.5.4l7.1-10.6h-7.2zm-11.5-.7l5.5.3l5.9-9.8h-6zm25.3 1.6l8 .5l8.7-11.6h-8.9zM6.2 22.8L3 28.9v2.8l3.2.2l5-9.1zm53.5 3.9L53.4 35l6.3.4z"/><path fill="#ebf2f2" d="m59.7 26.7l2.3-3.9V35l-2.3.4z"/><path fill="#686b6d" d="M59.7 22.8v3.9l2.3-3.9z"/><path fill="#333" d="M3 31.7v18L59.7 62V35.4z"/><g fill="#fff"><path d="M59.7 36.9L3 32.7v-1.2l56.7 3.7zm0 16.2L3 43.7v-1l56.7 8.9zm0-7.4L3 38.7v-1l56.7 6.6z"/><path d="m13.8 45l-1-.2v-5.4l1 .1z"/></g><path fill="#181919" d="m13.4 32.4l-11-.7V18.3z"/><path fill="#fff" d="M4.6 22.2c0 .6-.4 1-.9 1s-.8-.4-.8-1s.4-1 .8-1c.5 0 .9.5.9 1m0 7.7c0 .6-.4 1-.9 1s-.8-.5-.8-1c0-.6.4-1 .8-1c.5 0 .9.5.9 1m5.6.3c0 .6-.4 1-.9 1s-.9-.5-.9-1.1s.4-1 .9-1c.4 0 .9.5.9 1.1"/><path fill="#686b6d" d="M62 60.5L59.7 62V36.9l2.3-.4z"/><path fill="#ebf2f2" d="m62 35l-2.3.4v1.5l2.3-.4zm0 8.7l-2.3.6v1.4l2.3-.6zm0 7l-2.3 1v1.4l2.3-1z"/>`),
		g.Group(children),
	)
}
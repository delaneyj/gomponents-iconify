package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PoliceCar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#D0CFCE" d="m64.8 43.8l-1.1-.6c-.4-.2-.6-.6-.5-1c.3-1.9.5-8.5-9.7-11.5c-.2-.1-.4-.1-.6-.1l-19.6.1c-.4 0-.8.1-1.1.3l-10.3 6.9c-.2.1-.4.2-.6.2c-1.9-.1-3.7.1-5.6.4c-5.4 1.1-7.6 4-8.4 5.5c-.2.3-.2.7-.2 1c.1 2.4-1.5 5.1.9 7.3l19.4-.1l20.4-.5l16.1-.2c.9-.1 2.4-1.4 2.8-2.2c1.7-2.7-1.7-5.4-1.9-5.5z"/><path fill="#9B9B9A" d="M17.3 46.2c-2.2 0-4 1.8-4 4s1.8 4 4 4s4-1.8 4-4s-1.8-4-4-4zm39.8 0c-2.2 0-4 1.8-4 4s1.8 4 4 4s4-1.8 4-4s-1.8-4-4-4z"/><path fill="#3F3F3F" d="M56.1 39.1v-4.3c0-.9-.8-1.7-1.7-1.7H33.2c-.1 0-.2 0-.2.1l-8 5.7c-.2.1-.2.4-.1.6c.1.1.2.2.3.2c5.6 0 27.2-.2 30.4-.1c.3 0 .5-.2.5-.5c0 .1 0 .1 0 0z"/><path fill="#FCEA2B" d="m8.9 40.3l4 1.6l-2.1 2.9l-5 .1z"/><path fill="#92D3F5" d="M34.9 29.2s-.1-2.1 2.4-2.2c2.4-.1 2.4 2.2 2.4 2.2"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-miterlimit="10" d="M34.9 29.2s-.1-2.1 2.4-2.2c2.4-.1 2.4 2.2 2.4 2.2"/><path d="m47.6 50.8l-20.6.4m-18.9 0c-.9-.2-1.7-.5-1.8-1c-.1-1-.3-3.8-.3-5.1c0-.5.1-1.1.4-1.5c1.1-2 4.8-6.8 14.9-6.4l10.3-6.9c.5-.3 1.1-.5 1.6-.5l19.6-.1c.3 0 .6 0 .9.1c2.2.6 11.7 4 10.4 12.6l1.1.6c.5.2.9.7 1 1.2c.4 1.4.3 2.9-.2 4.3"/><path d="m24.9 39.7l30.2-.2V35"/><circle cx="17.3" cy="50.2" r="5"/><circle cx="57.1" cy="50.2" r="5"/><path d="m12.3 42.7l-1.8 2.2l-3.7-.1"/><path stroke-miterlimit="10" d="M24.1 43.2h26"/></g><path fill="none" stroke="#61B2E4" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2.1" d="M24.1 43.2h26"/>`),
		g.Group(children),
	)
}
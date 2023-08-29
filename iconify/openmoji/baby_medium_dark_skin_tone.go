package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BabyMediumDarkSkinTone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#92D3F5" d="M17 62.5v-1.9c0-8.4 1.1-25.7 4.1-32.1h30c3 6.3 3.9 23.6 3.9 32.1v1.9H17z"/><path fill="#61B2E4" d="M55 62.5H43V46c0-3-1-4-1-4l5.4-2.9l4.4-8l1.4 3.6L55 62.5z"/><circle cx="36" cy="35.5" r="3" fill="#EA5A47"/><path fill="#a57939" d="M51.1 25.7c0-9.1-6.2-15.7-15-15.7s-15 6.5-15 15.7c0 6.3-.8 12.2 7.1 16.3c0 0 3.5 1 7.9 1c5 0 8-1.1 8-1.1c8-4 7-10 7-16.2zM36 38.5c-1.7 0-3-1.3-3-3s1.3-3 3-3s3 1.3 3 3s-1.3 3-3 3z"/><g fill="none" stroke="#000" stroke-width="2"><path stroke-linecap="round" stroke-miterlimit="10" d="M27 28s1-1 2-1s2 1 2 1m10 0s1-1 2-1s2 1 2 1"/><path stroke-linecap="round" stroke-linejoin="round" d="M51 29.3c3 3 4 21.7 4 29.7v2.5m-38 0V59c0-8 1-26.7 4-29.7"/><path stroke-linecap="round" stroke-miterlimit="10" d="m32 16l-1 1m5-2l-1 2m5-1l1 1"/><circle cx="36" cy="35.5" r="3" stroke-miterlimit="10"/><path stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" d="M32.7 35.6c-1.1-.3-1.7-.6-1.7-.6m10 0s-.5.3-1.5.6"/><path stroke-linecap="round" stroke-linejoin="round" d="M44 42c8-3 7-12 7-15.8C51 17.3 44.8 11 36 11s-15 6.3-15 15.2C21 30 20 39 28.1 42c3.7 1.4 12.1 1.3 15.9 0z"/></g>`),
		g.Group(children),
	)
}
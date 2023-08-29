package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceWithThermometer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiFaceWithThermometer0" fill="none" stroke="#000" stroke-miterlimit="10" d="m40 49.3l-.7-.8c-.7-.9-.6-2.2.2-2.9l17-14.4c.9-.7 2.2-.6 2.9.2l.7.8c.7.9.6 2.2-.2 2.9l-17 14.4c-.8.8-2.1.7-2.9-.2z"/></defs><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M47.744 38.641c.1-.1.3-.2.5-.3c2-.6 3.9 1.4 3.3 3.4c-.1.3-.2.6-.5.8"/><circle cx="36" cy="36" r="24" fill="#FCEA2B"/><path fill="#FFF" d="m40 49.2l-.7-.8c-.7-.9-.6-2.2.2-2.9l17-14.4c.9-.7 2.2-.6 2.9.2l.7.8c.7.9.6 2.2-.2 2.9l-17 14.4c-.8.8-2.1.7-2.9-.2z"/><path fill="#EA5A47" d="m39.8 49l-.2-.3c-.6-.7-.5-1.7.2-2.3l8.6-7.3c.7-.6 1.7-.5 2.3.2l.2.3c.6.7.5 1.7-.2 2.3l-8.6 7.3c-.7.6-1.7.5-2.3-.2z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M50.975 27.2c-1.6-2.8-4.8-4.1-7.9-3.4"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M56.991 26.66C53.376 18.624 45.341 13 36 13c-12.655 0-23 10.245-23 23c0 12.655 10.245 23 23 23c10.646 0 19.585-7.231 22.197-16.974"/><path d="M30 33a3.001 3.001 0 0 1-6 0c0-1.655 1.345-3 3-3s3 1.345 3 3m18 0a3.001 3.001 0 0 1-6 0c0-1.655 1.345-3 3-3s3 1.345 3 3"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M21.025 27.2c.7-1.4 1.9-2.4 3.4-3c1.4-.7 3-.8 4.5-.4"/><use href="#openmojiFaceWithThermometer0" stroke-miterlimit="10"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M29.6 50.7c4.3-2.3 9.1-2.4 12.8 0"/><use href="#openmojiFaceWithThermometer0" stroke-miterlimit="10" stroke-width="2"/>`),
		g.Group(children),
	)
}
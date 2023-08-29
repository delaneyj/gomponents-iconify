package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mosque(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d0cfce" d="M54.975 35h-26a1 1 0 0 0-1 1v5.42h-5.5V16a1 1 0 0 0-1-1h-5a1 1 0 0 0-1 1v40a1 1 0 0 0 1 1h38.5a1 1 0 0 0 1-1V36a1 1 0 0 0-1-1Z"/><circle cx="41.967" cy="20.008" r="2" fill="#f4aa41"/><path fill="#9b9b9a" d="m41.975 41l-3.497 5.244a3 3 0 0 0-.503 1.664V56h8v-8.092a3 3 0 0 0-.504-1.664ZM21.908 28.192H16a1 1 0 0 1-.948-.681l-1.975-5.877a1.001 1.001 0 0 1 .948-1.318h9.899a1 1 0 0 1 .946 1.324l-2.016 5.877a1 1 0 0 1-.946.675Zm0 14.213H16a1 1 0 0 1-.948-.681l-1.975-5.877a1.001 1.001 0 0 1 .948-1.318h9.899a1 1 0 0 1 .946 1.324l-2.015 5.877a1 1 0 0 1-.947.675Z"/><path fill="#f4aa41" d="M53.435 37h-22.92a1.002 1.002 0 0 1-.877-.519A9.347 9.347 0 0 1 28.475 32c0-6.065 6.055-11 13.5-11s13.5 4.935 13.5 11a9.347 9.347 0 0 1-1.163 4.481a1.001 1.001 0 0 1-.877.519Zm-31.96-20c-.428 0-4.765.118-5.294-.044a1 1 0 0 1-.662-1.25l2.5-8.137a1 1 0 0 1 1.912 0l2.5 8.137a1 1 0 0 1-.662 1.25a1.013 1.013 0 0 1-.294.044Z"/><g fill="none" stroke="#000"><rect height="3" x="41.967" y="15.008" stroke-miterlimit="10"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M54.475 32a8.28 8.28 0 0 1-1.04 4h-22.92a8.28 8.28 0 0 1-1.04-4c0-5.52 5.6-10 12.5-10s12.5 4.48 12.5 10Zm.5 24h-38.5m12.5-3.076V36h26v16.924m-33.5-11.504h7.5"/><circle cx="41.967" cy="20.008" r="2" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><rect height="3" x="41.967" y="15.008" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M45.975 47.91V56h-8v-8.09a3.11 3.11 0 0 1 .5-1.67l2.84-4.25l.66-.99l.66.99l2.84 4.25a3.11 3.11 0 0 1 .5 1.67Zm-29.5 5.014V41.42m5 0v11.504m-5-31.608V16h5v5.316m.433 5.877H16l-1.975-5.877h9.899l-2.016 5.877zm0 14.212H16l-1.975-5.876h9.899l-2.016 5.876zM16.475 16l2.5-8.137l2.5 8.137m-5 11.193v8.336m5-8.336v8.336"/></g>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoPlayerGames(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor"><path stroke-linejoin="round" d="m20.584 11.426l.8-1.057l3.376-1.286l1.467 1.704l3.982 1.014l.123-1.576l2.104.64v3.957l-2.104.64l-.123-1.577l-3.982 1.015l-1.466 1.704l-3.378-1.286l-.8-1.057v-2.835Z"/><rect width="1.81" height="7.518" x="17.137" y="9.085" stroke-linecap="round" stroke-linejoin="bevel" rx=".363" ry=".363"/><path stroke-linecap="round" stroke-linejoin="round" d="M18.952 9.447V6.913c0-.151.121-.273.272-.273h3.944c.151 0 .273.122.273.273v2.673m7.319.769V8.457a.227.227 0 0 0-.227-.226h-3.269a.226.226 0 0 0-.226.226v2.537m-8.086 5.246v2.534c0 .151.121.273.272.273h3.944a.272.272 0 0 0 .273-.273v-2.7m7.319-.742v1.898a.227.227 0 0 1-.227.226h-3.269a.226.226 0 0 1-.226-.226v-2.537"/><path stroke-linejoin="round" d="m32.211 36.534l-.8 1.057l-3.377 1.286l-1.466-1.704l-3.982-1.014l-.123 1.576l-2.104-.64v-3.957l2.104-.64l.123 1.577l3.982-1.015l1.466-1.703l3.378 1.285l.799 1.057v2.835Z"/><rect width="1.81" height="7.518" x="33.84" y="31.358" stroke-linecap="round" stroke-linejoin="bevel" rx=".363" ry=".363"/><path stroke-linecap="round" stroke-linejoin="round" d="m33.84 38.513l.003 2.534a.272.272 0 0 1-.273.273h-3.943a.272.272 0 0 1-.273-.273v-2.672m-7.318-.635v1.763c0 .126.1.227.226.227h3.269a.226.226 0 0 0 .226-.227v-2.536m8.083-5.247l.003-2.534a.272.272 0 0 0-.273-.273h-3.943a.272.272 0 0 0-.273.273v2.673m-7.318.769V30.73c0-.125.1-.226.226-.226h3.269c.125 0 .226.101.226.226v2.537M8.822 10.13a3.21 3.21 0 0 1 3.842-3.144c1.345.258 2.419 1.421 2.553 2.784c.1 1.013-.222 2.013-.922 2.628c-1.297 1.138-5.473 4.206-5.473 4.206h6.412"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m33.84 30.972l5.272-20.473a2.087 2.087 0 0 0-1.501-2.542l-13.172-3.39a2.087 2.087 0 0 0-2.541 1.5l-.148.573m-2.811 10.92l-5.134 19.941a2.087 2.087 0 0 0 1.5 2.542l13.173 3.39a2.087 2.087 0 0 0 2.54-1.5l.159-.613"/>`),
		g.Group(children),
	)
}
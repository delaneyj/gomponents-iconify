package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trowel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d0cfce" d="M36 35.34v24.64s-7.518-6.788-8.104-8.13s-.277-16.26-.277-16.26z"/><path fill="#9b9b9a" d="M44.36 35.78L44.467 51s-6.624 7.454-8.463 8.976c-.022-2.034 0-23.9 0-23.9z"/><path fill="#a57939" d="M41.14 19.39c0-4.109-2.298-7.439-5.134-7.439s-5.134 3.331-5.134 7.439h.003l1.608 8.75H31.08v3.115h9.672V28.14h-.993z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M41.07 19.57c0-4.163-2.359-7.537-5.269-7.537s-5.269 3.375-5.269 7.537h.003l1.65 8.866h-1.438v3.156h9.927v-3.156h-1.02z"/><path d="M32.46 31.59v4.528h6.932V31.59H32.46z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m43.83 51.9l-.187-15.78h-15.48v15.86l7.832 8.001z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m32.38 36.12l2.281 5.303c.468 1.149 2.091 1.216 2.654.109l2.296-5.412"/><ellipse cx="35.8" cy="17.76" rx="1.915" ry="1.89"/>`),
		g.Group(children),
	)
}
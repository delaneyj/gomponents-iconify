package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coindcx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 19.6v20.9a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v12.1m24.304 4.6L36.8 34.76m0-10.56l-6.996 10.56"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.2 34.76V24.2h2.376a4.62 4.62 0 0 1 4.62 4.62v1.32a4.62 4.62 0 0 1-4.62 4.62Z"/><rect width="5.69" height="7.539" x="19.31" y="13.347" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.845"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.055 13.24v7.54m-11.54-1.327a2.844 2.844 0 0 1-2.47 1.433h0a2.845 2.845 0 0 1-2.845-2.845v-1.849a2.845 2.845 0 0 1 2.845-2.845h0a2.844 2.844 0 0 1 2.467 1.428M36.8 20.886v-4.694a2.845 2.845 0 0 0-2.845-2.845h0a2.845 2.845 0 0 0-2.845 2.845v4.694m0-4.694v-2.845m-3.168 17.872v.043a3.498 3.498 0 0 1-3.497 3.498h0a3.498 3.498 0 0 1-3.498-3.498v-3.564a3.498 3.498 0 0 1 3.498-3.498h0a3.498 3.498 0 0 1 3.498 3.498v.043"/><circle cx="28.055" cy="10.873" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}
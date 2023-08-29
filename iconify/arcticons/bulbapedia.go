package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bulbapedia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.219 6.546c.94 2.295 2.827 2.905 4.144 3.596"/><circle cx="24.222" cy="23.945" r="7.263" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.265 20.405c-2.581-6.552 1.79-10.243 6.838-10.788c5.164-.962 9.727.872 14-.427c1.93-.57 3.36-1.746 4.985-2.703c.799.539 1.777.798 2.628.926c-.042.24-1.914 1.85-3.177 2.667c-1.525.987-2.448 1.216-5.8 1.44c-4.387.017-9.257 1.464-11.807 5.66c-3.08 5.067.685 12.749.685 12.749c-3.062-2.657-6.705-5.821-8.352-9.523ZM22.658 36.71c-3.376-.858-7.333-2.996-9.04-6.781m9.04 6.781s13.59 3.978 17.104-6.741c.774-3.306-.96-6.463-2.405-9.316c-1.342-2.306-1.605-5.136-.794-7.666c.192-.663.89-2.26 1.091-2.157c-.367 2.368.104 4.785 1.147 6.928c1.54 3.467 3.548 6.798 4.317 10.551c1.108 5.005-.081 10.458-5.038 12.36c-1.845.898-3.958 1.052-5.946.594c-3.426-.895-6.888-2.887-9.476-4.554Zm-5.699-12.765h-5.131m19.657 0h7.454"/>`),
		g.Group(children),
	)
}
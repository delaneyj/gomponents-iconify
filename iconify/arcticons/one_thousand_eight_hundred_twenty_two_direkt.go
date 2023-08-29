package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OneThousandEightHundredTwentyTwoDirekt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.248 31.136c0-1.22.993-2.216 2.208-2.216h0m-2.208-.001v5.874m9.655-8.866v8.866m0-1.884l3.975-3.99m-2.761 2.771l3.202 3.103"/><circle cx="17.332" cy="26.26" r=".747" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.332 28.919v5.874m18.564-7.758v7.758m-1.215-5.874H37m-10.362 4.766c-.331.665-1.104 1.108-1.877 1.108h0a2.219 2.219 0 0 1-2.208-2.216v-1.44c0-1.22.993-2.217 2.208-2.217h0c1.214 0 2.208.997 2.208 2.216v.776h-4.416m-7.137-.776c0-1.22-.994-2.216-2.208-2.216h0A2.219 2.219 0 0 0 11 31.136v1.44c0 1.22.994 2.217 2.208 2.217h0a2.219 2.219 0 0 0 2.208-2.216m0 2.216v-8.866m-5.588-3.804h4.337m-4.337-7.595l2.169-1.211m0 0v8.806m6.807-4.403c-1.192 0-2.168.99-2.168 2.201h0c0 1.211.976 2.202 2.168 2.202h1.41c1.193 0 2.169-.99 2.169-2.202h0c0-1.21-.976-2.201-2.169-2.201m0 0c1.193 0 2.169-.99 2.169-2.202h0c0-1.21-.976-2.201-2.169-2.201h-1.41c-1.193 0-2.169.99-2.169 2.201h0c0 1.211.976 2.202 2.17 2.202h1.409Zm4.642-1.541c0-1.651 1.302-2.972 2.82-2.972s2.928 1.32 2.928 2.972a3.03 3.03 0 0 1-.868 2.091c-1.193.99-4.88 3.853-4.88 3.853h5.748m2.476-5.944c0-1.651 1.302-2.972 2.82-2.972s2.928 1.32 2.928 2.972a3.03 3.03 0 0 1-.868 2.091c-1.192.99-4.88 3.853-4.88 3.853h5.748"/>`),
		g.Group(children),
	)
}
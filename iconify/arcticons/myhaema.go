package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Myhaema(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2ZM8.531 18.496v7.008m4.643-7.008v7.008m-4.643-3.517h4.643"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.253 22.7a1.752 1.752 0 0 1 1.752-1.751h0a1.752 1.752 0 0 1 1.752 1.752v2.803m-3.504-4.555v4.555"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.756 22.7a1.752 1.752 0 0 1 1.752-1.751h0A1.752 1.752 0 0 1 32.26 22.7v2.803m-8.97-.883a1.751 1.751 0 0 1-1.522.884h0a1.752 1.752 0 0 1-1.752-1.752v-1.139a1.752 1.752 0 0 1 1.752-1.752h0a1.752 1.752 0 0 1 1.752 1.752v.57h-3.504m-1.669.569a1.752 1.752 0 0 1-1.752 1.752h0a1.752 1.752 0 0 1-1.752-1.752v-1.139a1.752 1.752 0 0 1 1.752-1.752h0a1.752 1.752 0 0 1 1.752 1.752m0 2.891v-4.643m19.185 2.891a1.752 1.752 0 0 1-1.752 1.752h0a1.752 1.752 0 0 1-1.753-1.752v-1.139a1.752 1.752 0 0 1 1.752-1.752h0a1.752 1.752 0 0 1 1.752 1.752m.001 2.891v-4.643m-12.279 6.861c.526 2.53 3.296 3.317 7.181 3.322c3.237.005 6.432-.4 7.182-3.764"/><circle cx="39.616" cy="25.136" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}
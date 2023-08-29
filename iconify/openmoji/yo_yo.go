package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YoYo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#8967aa" d="M18.917 32.873a13.33 13.33 0 1 1 17.706 19.259"/><circle cx="24.328" cy="45.662" r="13.328" fill="#b399c8"/><path fill="#fcea2b" d="m24.969 54.8l-3.234-4.978l-5.934-.157l3.734-4.614l-1.684-5.692l5.542 2.126l4.893-3.361l-.309 5.928l4.708 3.615l-5.733 1.538l-1.983 5.595z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round"><path stroke-width="2" d="m24.969 54.8l-3.234-4.978l-5.934-.157l3.734-4.614l-1.684-5.692l5.542 2.126l4.893-3.361l-.309 5.928l4.708 3.615l-5.733 1.538l-1.983 5.595z"/><circle cx="24.328" cy="45.662" r="13.328" stroke-width="2"/><path stroke-width="2" d="M18.917 32.873a13.33 13.33 0 1 1 17.706 19.259"/><path stroke-width="2" d="M53.594 12.795c1.276 3.06 5.58 2.747 5.58 2.747a3.18 3.18 0 0 0 3.502-3.104a4.168 4.168 0 0 0-4.413-1.911s-10.068.39-12.281 13.441c-1.513 8.919-4.51 9.777-8.863 9.316c-1.62-.171-2.912-.994-5.784-2.265a9.93 9.93 0 0 0-2.54-.724s-3.395-.297-4.467.457"/><path stroke-width=".5" d="M16.57 35.01a13.33 13.33 0 1 1 17.706 19.26"/><path stroke-width=".5" d="M15.801 35.541A13.33 13.33 0 1 1 33.507 54.8"/><path stroke-width=".5" d="M15.007 36.335a13.33 13.33 0 1 1 17.707 19.258"/></g>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VtvGo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m5.593 9.264l7.427 13.023c.767 1.33 2.296 1.33 3.064 0l4.595-8.426v9.447h7.149V13.86l4.085 8.425c.767 1.33 2.297 1.329 3.064 0L42.4 9.268c.297-.41-.018-.768-.531-.768h-4.596c-.515 0-1.235.362-1.493.8l-2.336 4.561l-2.335-4.557c-.257-.445-.983-.804-1.495-.804H18.38c-.513 0-1.18.316-1.407.771l-2.423 4.59l-2.416-4.575c-.172-.476-.903-.786-1.414-.786H6.126c-.513-.01-.807.334-.534.763h0Zm22.235 4.598h5.617m-12.766 0L23.774 8.5m.224 28.243a2.043 2.043 0 0 1-2.042-2.042v-1.328a2.042 2.042 0 1 1 4.085 0V34.7a2.042 2.042 0 0 1-2.043 2.042Zm-3.983-5.413v6.127a2.042 2.042 0 0 1-3.487 1.445"/><path d="M17.973 31.33c1.128 0 2.042.914 2.042 2.042V34.7a2.042 2.042 0 1 1-4.085 0v-1.328c0-1.128.915-2.042 2.043-2.042Z"/></g><circle cx="28.236" cy="28.828" r=".766" fill="currentColor"/><circle cx="30.023" cy="28.828" r=".766" fill="currentColor"/><circle cx="31.81" cy="28.828" r=".766" fill="currentColor"/><circle cx="30.023" cy="30.615" r=".766" fill="currentColor"/><circle cx="31.81" cy="30.615" r=".766" fill="currentColor"/><circle cx="31.81" cy="32.403" r=".766" fill="currentColor"/>`),
		g.Group(children),
	)
}
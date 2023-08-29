package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M24 6c-4.5 0-7 1.2-7 1.2V12l-3.529 3.529c-.593.593-.236 1.588.6 1.648c2.017.143 5.434.323 9.929.323c2.206 0 4.152-.043 5.8-.104h-.017a6 6 0 1 1-11.567 0c-.74-.027-1.42-.058-2.036-.09a8 8 0 1 0 15.64 0a112.94 112.94 0 0 0 2.109-.13c.836-.06 1.193-1.054.6-1.647L30.999 12V7.2S28.5 6 24 6Zm-5 6.828l-2.492 2.492c1.93.097 4.462.18 7.492.18c3.03 0 5.562-.083 7.492-.18L29 12.828V8.62a13.77 13.77 0 0 0-.302-.08C27.656 8.276 26.07 8 24 8c-2.071 0-3.656.276-4.698.54a8.595 8.595 0 0 0-.302.08v4.208ZM30.148 9.01Zm0 0l-.002-.002l.001.001Z" clip-rule="evenodd"/><path d="m25 30.341l5 1.488V40h-5v-9.659ZM23 40h-5v-8.17l5-1.489V40Z"/></g>`),
		g.Group(children),
	)
}
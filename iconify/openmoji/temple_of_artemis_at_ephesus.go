package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TempleOfArtemisAtEphesus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#9b9b9a" d="M17.508 48.797a1 1 0 0 1-1-1v-15.75a1 1 0 0 1 2 0v15.75a1 1 0 0 1-1 1Zm8 0a1 1 0 0 1-1-1v-15.75a1 1 0 0 1 2 0v15.75a1 1 0 0 1-1 1Zm8.001 0a1 1 0 0 1-1-1v-15.75a1 1 0 0 1 2 0v15.75a1 1 0 0 1-1 1Zm8 0a1 1 0 0 1-1-1v-15.75a1 1 0 0 1 2 0v15.75a1 1 0 0 1-1 1Zm8 0a1 1 0 0 1-1-1v-15.75a1 1 0 0 1 2 0v15.75a1 1 0 0 1-1 1Zm8 0a1 1 0 0 1-1-1v-15.75a1 1 0 0 1 2 0v15.75a1 1 0 0 1-1 1Z"/><path fill="#d0cfce" d="M12.946 28.07h46.109v3.978H12.946z"/><rect width="48.109" height="5.582" x="11.946" y="46.797" fill="#d0cfce" rx="1"/><rect width="55.674" height="5.581" x="8.163" y="50.379" fill="#d0cfce" rx="1"/><path fill="#d0cfce" d="M59.054 28.07H12.946L36 20.705l23.054 7.365z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M59.054 47.797H12.946m49.891 3.582H9.163M66.29 55H5.71M16 32.048v15.749m8-15.749v15.749m8-15.749v15.749m8-15.749v15.749m8-15.749v15.749m8-15.749v15.749m3.054-19.727H12.946L36 20.705l23.054 7.365z"/><path d="M59.054 25.155v6.893H12.946v-6.893M36 20.705v-2.319"/></g>`),
		g.Group(children),
	)
}
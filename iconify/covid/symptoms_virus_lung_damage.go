package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SymptomsVirusLungDamage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18.058 9.159a2.967 2.967 0 1 0 0-5.934a2.967 2.967 0 0 0 0 5.934ZM17.563 1h.989m-.494 0v2.225m3.322-1.054l.699.699m-.35-.349l-1.573 1.573m3.094 1.604v.989m0-.495h-2.225m1.054 3.322l-.699.699m.349-.349L20.156 8.29m-1.604 3.095h-.989m.495 0V9.159m-3.322 1.054l-.699-.699m.349.35L15.96 8.29m-3.095-1.603v-.989m0 .494h2.226M14.037 2.87l.699-.699m-.35.35l1.574 1.573m-6.062.609v7.929M6.849 9.078a1.219 1.219 0 0 0-2.122-.82l-.805.884A12.2 12.2 0 0 0 .75 17.35v3.21a2.44 2.44 0 0 0 3.35 2.266l1.986-.8a1.22 1.22 0 0 0 .767-1.133L6.849 9.078Zm-2.44 7.823l5.489-4.269m8.834 1.961c.21.903.315 1.828.315 2.755v3.212a2.44 2.44 0 0 1-3.347 2.266l-1.986-.8a1.22 1.22 0 0 1-.767-1.133v-9.267m2.44 5.275l-5.489-4.269"/>`),
		g.Group(children),
	)
}
package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SyncOccurence(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 256h128v512h-512V640h281q-51-88-122-159t-158-121t-184-77t-201-27q-138 0-263 46T536 431T366 628t-96 249l-126-25q20-105 63-199t106-174t141-146t170-110t192-70t208-25q118 0 230 30t211 87t183 137t144 180V256zm-768 1536q138 0 263-46t225-129t170-196t96-249l125 24q-30 158-111 290t-198 229t-263 151t-307 54q-119 0-231-30t-212-87t-182-136t-143-181v306H128v-512h512v128H359q51 89 122 160t158 120t183 77t202 27z"/>`),
		g.Group(children),
	)
}
package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fish(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M15.38 32.844C9.725 31.429.532 34.965 9.725 38.501c3.535 9.192 7.07 0 5.657-5.657Zm7.102-22.596c-4.903-.014-11.232 1.536-12.029 3.727c-.577 1.589 2.425 3.702 6.25 4.818m21.275 6.949c.013 4.904-1.537 11.232-3.728 12.03c-1.589.577-3.702-2.425-4.818-6.251"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M38 25.765C44.662 19.7 40.124 8.09 40.124 8.09S28.108 4.56 22.45 10.216c-5.657 5.657-7.071 22.628-7.071 22.628s15.956-1.015 22.62-7.08Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M38.008 25.773s-6.718-1.06-10.607-4.95c-3.89-3.889-4.95-10.606-4.95-10.606"/><circle cx="33.766" cy="14.459" r="2" fill="currentColor" transform="rotate(45 33.766 14.46)"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M41 20.945c-.54 1.76-1.478 3.434-3 4.82c-1.648 1.5-3.864 2.69-6.266 3.633M27 7.656c-1.697.489-3.282 1.294-4.549 2.56c-1.184 1.186-2.183 2.867-3.022 4.784"/></g>`),
		g.Group(children),
	)
}
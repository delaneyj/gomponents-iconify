package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BsnlSelfcare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.344 6.726c3.067-1.387 5.422-1.607 6.524-.61c2.248 2.031-1.22 8.46-7.747 14.358q-.46.416-.931.82m-2.941-2.316l3.959 3.959l-3.976.017Zm-5.593 9.04c-3.067 1.387-5.422 1.607-6.524.61c-2.248-2.031 1.22-8.46 7.747-14.358h0q.46-.416.931-.82m2.867 2.09l-4.355-4.355l4.374-.019Z"/><ellipse cx="23.664" cy="17.305" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="9.568" ry="9.902"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.06 37.542a2.479 2.479 0 0 1 0 4.958H7.97v-9.915h4.09a2.479 2.479 0 1 1 0 4.957Zm23.012-5.012v9.915h4.957m-14.288.014v-9.916l6.569 9.916v-9.916M16.903 41.34a2.955 2.955 0 0 0 2.479 1.115h1.487a2.486 2.486 0 0 0 2.479-2.478h0a2.486 2.486 0 0 0-2.479-2.48h-1.611a2.486 2.486 0 0 1-2.479-2.478h0a2.486 2.486 0 0 1 2.479-2.479h1.487a2.661 2.661 0 0 1 2.479 1.115m-11.163 3.887h-4.09"/>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlQuranIndonesia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.224 13.542c2.496-2.86 4.767-5.13 3.701-6.28S8.613 12.515 9.586 14.187c.58.996 15.989 12.178 17.509 13.078c1.785 1.056 4.653-.18 4.4-2.144L29.168 7.11s1.178 1.582 1.542 2.675m4.374 16.722c-.252-1.963-2.327-18.012-2.327-18.012s1.178 1.582 1.542 2.676"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.383 33.87c0-6.218-2.09-20.011-2.09-20.011s1.178 1.581 1.543 2.675m-5.766-1.987c-.364-1.094-1.541-2.676-1.541-2.676s.593 7.208 4.041 8.638c-9.246-4.479-13.647 10.178-17.642 5.888c-2.851-3.06 2.27-7.36 2.27-7.36"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.916 14.03c-.365-1.093-1.542-2.675-1.542-2.675S38.4 29.25 38.832 33.87c.205 2.201-3.874 3.21-5.388-.49s-.126-4.753.799-4.753s2.44 3.084 0 4.514s-6.505 2.468-7.738 3.897s-3.393 4.71-4.823 4.71s-4.233-1.541-1.037-7.878m20.804-2.72c-.253-3.28-1.963-16.57-1.963-16.57s1.178 1.581 1.542 2.675M13.993 8.944c1.724-1.073 5.278-1.745 7.128-3.196M8.981 27.193C6.682 29.72 7.327 34.85 10.02 34.01c1.005-.314 2.215-2.404 4.233-3.308c3.387-1.516 7.719-2.29 7.458-4.71c-.336-3.126-3.336-4.945-3.336-4.945"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.818 29.019c-1.437-.635-1.907-1.367-1.907-1.367c1.164 2.6 2.061 4.184 2.09 5.432s-6.407 5.236-9.17 5.236c-3.216 0-5.57-2.312-4.656-6.15m-2.04-17.435c-2.114.98-3.425 5.617-1.813 5.149"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}
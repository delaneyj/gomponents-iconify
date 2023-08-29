package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandGoogleMaps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9.5 9.5a2.5 2.5 0 1 0 5 0a2.5 2.5 0 1 0-5 0m-3.072 2.994l7.314-9.252m-3.74 4.693L7.065 5.39m10.628 1.203l-8.336 9.979"/><path d="M17.591 6.376c.472.907.715 1.914.709 2.935a7.263 7.263 0 0 1-.72 3.18a19.085 19.085 0 0 1-2.089 3c-.784.933-1.49 1.93-2.11 2.98c-.314.62-.568 1.27-.757 1.938c-.121.36-.277.591-.622.591c-.315 0-.463-.136-.626-.593a10.595 10.595 0 0 0-.779-1.978a18.18 18.18 0 0 0-1.423-2.091c-.877-1.184-2.179-2.535-2.853-4.071A7.077 7.077 0 0 1 5.7 9.3a6.226 6.226 0 0 1 1.476-4.055A6.25 6.25 0 0 1 11.987 3a6.462 6.462 0 0 1 1.918.284a6.255 6.255 0 0 1 3.686 3.092z"/></g>`),
		g.Group(children),
	)
}
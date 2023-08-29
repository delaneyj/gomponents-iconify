package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HourglassCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopHourglassCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M10.203 7c.13.819.323 1.595.575 2.084c.198.385.586.874 1.118 1.407c.365.365.752.706 1.104.996c.352-.29.74-.631 1.104-.996c.532-.533.92-1.022 1.118-1.407c.252-.489.444-1.265.575-2.084h-5.594Zm-.662-2c-.844 0-1.518.697-1.42 1.536c.125 1.076.378 2.49.879 3.464c.672 1.305 2.218 2.643 3.18 3.393c.485.38 1.155.38 1.64 0c.962-.75 2.508-2.088 3.18-3.393c.501-.973.754-2.388.88-3.464c.097-.84-.577-1.536-1.421-1.536H9.54Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M15.797 19c-.13-.819-.323-1.595-.575-2.084c-.198-.384-.586-.875-1.118-1.407A15.7 15.7 0 0 0 13 14.513c-.352.29-.74.631-1.104.996c-.532.532-.92 1.023-1.118 1.407c-.252.489-.444 1.265-.575 2.084h5.594Zm.662 2c.844 0 1.518-.697 1.42-1.535c-.125-1.077-.378-2.492-.879-3.465c-.672-1.305-2.218-2.643-3.18-3.393a1.326 1.326 0 0 0-1.64 0c-.962.75-2.508 2.088-3.18 3.393c-.501.973-.754 2.388-.88 3.465c-.097.838.577 1.535 1.421 1.535h6.918Z" clip-rule="evenodd"/><path d="M10 18.75s2-1.5 3-1.5s3 1.5 3 1.5v.5h-6v-.5Z"/><path fill-rule="evenodd" d="M7 5.5a1 1 0 0 1 1-1h10a1 1 0 1 1 0 2H8a1 1 0 0 1-1-1Zm0 15a1 1 0 0 1 1-1h10a1 1 0 1 1 0 2H8a1 1 0 0 1-1-1Z" clip-rule="evenodd"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopHourglassCircleFilled0)"/></g>`),
		g.Group(children),
	)
}
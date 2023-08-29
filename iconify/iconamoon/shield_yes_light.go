package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldYesLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="m12 3l.295-.69a.75.75 0 0 0-.59 0L12 3Zm0 18l-.372.651a.75.75 0 0 0 .744 0L12 21Zm6.394-15.26l-.295.69l.295-.69ZM8.024 18.727l-.373.652l.372-.652Zm3.68-16.416L5.312 5.05L5.9 6.43l6.394-2.74l-.59-1.38ZM4.25 6.659v6.86h1.5v-6.86h-1.5Zm3.401 12.72l3.977 2.272l.744-1.302l-3.977-2.273l-.744 1.303Zm4.721 2.272l3.977-2.272l-.744-1.303l-3.977 2.273l.744 1.302Zm7.378-8.133V6.66h-1.5v6.86h1.5Zm-1.06-8.467l-6.395-2.74l-.59 1.378l6.394 2.74l.59-1.378Zm1.06 1.608c0-.7-.417-1.332-1.06-1.608l-.591 1.379a.25.25 0 0 1 .151.23h1.5Zm-3.401 12.72a6.75 6.75 0 0 0 3.401-5.86h-1.5a5.25 5.25 0 0 1-2.645 4.557l.744 1.303ZM4.25 13.519a6.75 6.75 0 0 0 3.401 5.86l.744-1.303a5.25 5.25 0 0 1-2.645-4.558h-1.5ZM5.31 5.05a1.75 1.75 0 0 0-1.06 1.608h1.5c0-.1.06-.19.152-.23L5.31 5.052Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m15 10l-4 4l-2-2"/></g>`),
		g.Group(children),
	)
}
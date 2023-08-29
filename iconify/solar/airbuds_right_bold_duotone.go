package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirbudsRightBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M19.618 22C20.933 22 22 20.88 22 19.5h-4.765c0 1.38 1.067 2.5 2.383 2.5Z"/><path fill-rule="evenodd" d="M13 7.889V5.542c0-.194 0-.29.004-.372c.08-1.713 1.385-3.082 3.017-3.166c.078-.004.17-.004.354-.004c.307 0 .46 0 .59.007c2.72.14 4.895 2.422 5.029 5.276c.006.136.006.297.006.62v9.93h-4.765v-5.5c0-.613-.474-1.11-1.058-1.11C14.422 11.222 13 9.73 13 7.888Zm1.677-2.764c0-.46.335-.833.75-.833c.414 0 .75.373.75.833v2.778c0 .46-.336.833-.75.833c-.415 0-.75-.373-.75-.833V5.125Z" clip-rule="evenodd"/><path d="M13 16.5a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0Z" opacity=".5"/><path fill-rule="evenodd" d="M10.735 4.95a.75.75 0 0 1-.586.885a4.256 4.256 0 0 0-3.314 3.314a.75.75 0 0 1-1.47-.298A5.756 5.756 0 0 1 9.85 4.365a.75.75 0 0 1 .884.586Z" clip-rule="evenodd" opacity=".5"/><path fill-rule="evenodd" d="M6 13.25a.75.75 0 0 0-.75.75v5a.75.75 0 0 0 1.5 0v-1.75h.325l1.282 2.136a.75.75 0 1 0 1.286-.772l-.99-1.649c.629-.33 1.097-.93 1.097-1.715c0-1.227-1.141-2-2.25-2H6Zm.75 1.5v1h.75c.548 0 .75-.346.75-.5c0-.154-.202-.5-.75-.5h-.75Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
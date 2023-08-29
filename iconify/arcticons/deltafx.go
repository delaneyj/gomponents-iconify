package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Deltafx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.389 29.596s-23.559 1.344-28.464.642"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.672 29.742c-.028.014-1.525.67 2.16 9.97m-.001-.001s-9.668 1.131-15.099 2.15c-8.802 1.653-14.371-.423-14.371-.423a11.878 11.878 0 0 1 1.67-7.925a17.53 17.53 0 0 1 2.019-3.117"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.524 34.482s-2.945-.032-2.878-4.41c.07-4.5 4.062-1.775 14.937-2.666c8.166-.67 19.477-.887 19.976 1.295c0 0 1.153 5.023-2.293 4.973s-3.256.808-3.256.808"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.141 28.202S4.86 14.382 6.378 10.87s8.068-1.122 11.519-1.799s19.58-2.527 21.84.438s2.838 3.151 2.11 12.275c-.227 5.989-.288 6.917-.288 6.917"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.323 9.694s-.987-3.972 2.286-4.009s6.504-1.334 5.851 3.493m8.986 13.344c-1.217.257-4.492.67-4.981-2.95c-.588-4.354 6.515-5.728 7.443-1.681c.8 3.49-1.748 4.48-2.462 4.631Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.805 13.06c1.297 1.039 3.501 2.651 3.5 6.102s-3.46 6.122-5.954 6.33c-2.413.202-7.339-1.034-7.514-6.065c-.142-4.107 1.582-6.106 4.696-6.954c2.232-.607 3.976-.452 5.272.587Zm13.631-2.64l-6.819.517a6.317 6.317 0 0 1 .27 2.64c-.18.97 8.066-.039 8.066-.039"/><path fill="none" stroke="currentColor" d="m21.944 30.332l-.21-2.94m5.626 0l.255 2.932M20.71 12.737l-.116-4.08m6.211 4.403l-.743-4.955"/><circle cx="10.786" cy="23.712" r="1.528" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="37.717" cy="23.712" r="1.528" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}
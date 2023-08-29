package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Portugal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#e04122" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M45 11.75h-6.32A45.89 45.89 0 0 1 24 9.38A45.73 45.73 0 0 0 9.37 7H3c-.58 0-1 .35-1 .79v26a.94.94 0 0 0 1 .79h6.37A46 46 0 0 1 24 37a45.62 45.62 0 0 0 14.65 2.38H45a.93.93 0 0 0 1-.79v-26a.94.94 0 0 0-1-.84Z"/><path fill="#46b000" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M16.19 7.51A46.8 46.8 0 0 0 9.37 7H3c-.58 0-1 .35-1 .79v26a.94.94 0 0 0 1 .79h6.37a46.8 46.8 0 0 1 6.82.5Z"/><path fill="#ffe500" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M9.609 25.168a8.33 7.55 60.71 1 0 13.17-7.387a8.33 7.55 60.71 1 0-13.17 7.387Z"/><path fill="#e04122" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M23.62 19.82a49.24 49.24 0 0 0-14.5-2.27m14.72 5.81a71.11 71.11 0 0 0-15.36-1.88m13.82 5.46a45.93 45.93 0 0 0-12.76-1.85"/><path fill="#e04122" stroke="#45413c" d="m19.48 17.47l-6.58-.57a.9.9 0 0 0-1 .91v3.63A6.16 6.16 0 0 0 15.64 27a1.71 1.71 0 0 0 1.11.09a5 5 0 0 0 3.73-4.9v-3.63a1.11 1.11 0 0 0-1-1.09Z"/><path fill="#fff" d="M18.38 18.92L14 18.53a.46.46 0 0 0-.5.46v2.72a3.86 3.86 0 0 0 2.34 3.48a1.07 1.07 0 0 0 .7.06a3.11 3.11 0 0 0 2.34-3.07v-2.72a.54.54 0 0 0-.5-.54Z"/><path fill="#009fd9" d="M16.1 22.75a.28.28 0 0 0 .18 0a.8.8 0 0 0 .6-.79v-.74l-1.38-.12v.74a1 1 0 0 0 .6.91Zm-1.67-.15a.25.25 0 0 0 .17 0a.8.8 0 0 0 .6-.79v-.74L13.83 21v.74a1 1 0 0 0 .6.86Zm3.35.3a.28.28 0 0 0 .18 0a.79.79 0 0 0 .6-.78v-.74l-1.38-.12V22a1 1 0 0 0 .6.9Zm-1.68-2.02a.28.28 0 0 0 .18 0a.79.79 0 0 0 .6-.78v-.75l-1.38-.12V20a1 1 0 0 0 .6.88Zm0 3.86a.28.28 0 0 0 .18 0a.8.8 0 0 0 .6-.79v-.74l-1.38-.12v.74a1 1 0 0 0 .6.91Z"/><path fill="#45413c" d="M12.5 45.5a11.5 1.5 0 1 0 23 0a11.5 1.5 0 1 0-23 0Z" opacity=".15"/>`),
		g.Group(children),
	)
}
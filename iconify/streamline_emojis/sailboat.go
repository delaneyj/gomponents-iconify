package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sailboat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M8.98 45.39a15.02 1.61 0 1 0 30.04 0a15.02 1.61 0 1 0-30.04 0Z" opacity=".15"/><path fill="#bf8256" d="M39.81 37.62a.88.88 0 0 1 .87 1c-.75 5.94-7.44 6.89-17.07 6.89s-16.33-1-17.08-6.89a.88.88 0 0 1 .87-1Z"/><path fill="#dea47a" d="M7 40.36a.92.92 0 0 1 .43-.11h32.38a.89.89 0 0 1 .43.11a7.12 7.12 0 0 0 .44-1.75a.88.88 0 0 0-.87-1H7.4a.88.88 0 0 0-.87 1A7.12 7.12 0 0 0 7 40.36Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M39.81 37.62a.88.88 0 0 1 .87 1c-.75 5.94-7.44 6.89-17.07 6.89s-16.33-1-17.08-6.89a.88.88 0 0 1 .87-1Z"/><path fill="#ff6242" d="M26 5.21h-1.1v28h15.55a.88.88 0 0 0 .65-1.47a67 67 0 0 1-14.25-25.9a.87.87 0 0 0-.85-.63Z"/><path fill="#ffe500" d="M29.4 13.08a35.69 35.69 0 0 1-4.48 1.67v5.95a35.72 35.72 0 0 0 6.82-2.63c-.83-1.62-1.63-3.27-2.34-4.99Zm4.71 9.23a35.27 35.27 0 0 1-9.19 3.54v5.89a34.18 34.18 0 0 0 12.13-5C36 25.28 35 23.82 34.11 22.31Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M26 5.21h-1.1v28h15.55a.88.88 0 0 0 .65-1.47a67 67 0 0 1-14.25-25.9a.87.87 0 0 0-.85-.63Z"/><path fill="#ff6242" d="M23.17 1.71V35H4.38a.86.86 0 0 1-.61-1.49c7.67-7.43 14.02-18.33 19.4-31.8Z"/><path fill="#ffe500" d="M23.17 11.5c-1.86.86-3.68 1.58-5.4 2.18q-2.34 4.55-4.88 8.49a63.11 63.11 0 0 0 10.28-3.66Zm-19.4 22a.86.86 0 0 0 .61 1.5h7.28a57.14 57.14 0 0 0 11.51-5.25v-7A59.82 59.82 0 0 1 7.8 29.14a58.56 58.56 0 0 1-4.03 4.36Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M23.17 1.71V35H4.38a.86.86 0 0 1-.61-1.49c7.67-7.43 14.02-18.33 19.4-31.8Z"/><path fill="#bf8256" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M23.17 1.71h1.75v35.91h-1.75z"/>`),
		g.Group(children),
	)
}
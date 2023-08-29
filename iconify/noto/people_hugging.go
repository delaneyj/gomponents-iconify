package noto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PeopleHugging(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<radialGradient id="notoPeopleHugging0" cx="71.353" cy="54.761" r="41.042" gradientTransform="rotate(3.865 250.515 86.799)" gradientUnits="userSpaceOnUse"><stop offset=".47" stop-color="#42A5F5"/><stop offset="1" stop-color="#FCE8B2"/></radialGradient><path fill="url(#notoPeopleHugging0)" d="M82.58 5.72c-10.59-.71-27.85 9.87-23.63 34.9c2.61 15.47 12.62 23.12 19.68 23.6s18.01-5.76 22.68-20.74c7.54-24.23-8.14-37.05-18.73-37.76z"/><linearGradient id="notoPeopleHugging1" x1="75.601" x2="66.573" y1="30.528" y2="114.104" gradientUnits="userSpaceOnUse"><stop offset=".128" stop-color="#FCE8B2"/><stop offset=".64" stop-color="#42A5F5"/></linearGradient><path fill="url(#notoPeopleHugging1)" d="M117.94 83.27c-4.15-16.61-16.84-25.96-30.58-27.66c-5.48-.68-11.04-.92-16.48.01c-1.27.22-2.7.82-3.97.86c-.98.03-32.26-12.01-33.57-12.46C29.68 42.73 22.17 43 19.17 51c-3.93 10.47 12 18 35 36l-.94 37h66.79s1.08-28.09-2.08-40.73z"/><linearGradient id="notoPeopleHugging2" x1="28.109" x2="63.83" y1="135.296" y2="46.551" gradientUnits="userSpaceOnUse"><stop offset="0" stop-color="#0D47A1"/><stop offset=".679" stop-color="#1976D2"/></linearGradient><path fill="url(#notoPeopleHugging2)" d="M89.17 70c-3-5-9.75-1.11-12 1c-16 15-19.7 17.19-22.96 15.43c-.99-.53-2.94-1.91-8.22-5.99S22.17 62 22.17 62c-18 15-13.96 62-13.96 62h45.02l.81-14.92c5.68.91 9.99-1.04 14.12-4.08c7.61-5.62 13.02-10.7 18-19c3.01-5 6.01-11 3.01-16z"/><radialGradient id="notoPeopleHugging3" cx="45.628" cy="49.205" r="46.585" gradientTransform="rotate(-22.626 44.006 45.518)" gradientUnits="userSpaceOnUse"><stop offset=".527" stop-color="#1565C0"/><stop offset="1" stop-color="#FADA80"/></radialGradient><path fill="url(#notoPeopleHugging3)" d="M39.45 7.91c-9.8 4.08-20.53 21.26-5.58 41.78c9.23 12.68 21.61 15.06 28.14 12.34c6.53-2.72 13.55-13.19 11.05-28.67C69 8.3 49.25 3.83 39.45 7.91z"/><path fill="#1976D2" d="M103.75 60.18S93.18 67 89.18 60c-3.58-6.26 7.1-12.75 7.1-12.75c3.11-1.79 7.23-1.56 9.56 2.53c2.79 4.89 1.02 8.61-2.09 10.4z"/>`),
		g.Group(children),
	)
}
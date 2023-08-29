package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RaspiCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="18.406" r="5.224" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.167 31.588q-.003.073-.003.146a4.005 4.005 0 0 0 5.53 3.701"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.184 23.863a4.129 4.129 0 0 0 3.93 7.261m21.258 5.685a3.41 3.41 0 0 0 2.686-6.046"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.733 37.82a3.873 3.873 0 1 0 6.653-2.81"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.983 35.546a5.114 5.114 0 0 0-.027.522a5.045 5.045 0 0 0 9.969 1.095"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.224 40.596a3.522 3.522 0 0 0 5.015.734M15.067 28.954A6.011 6.011 0 0 0 21 35.894a5.996 5.996 0 0 0 4.81-2.41m10.314-1.336a5.225 5.225 0 0 0 1.538-9.928"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m32.53 26.017l.106.001a5.224 5.224 0 0 0 3.089-9.437M25.92 30.758a5.225 5.225 0 1 0 6.227-3.504M18.85 19.285A5.224 5.224 0 1 0 21 26.397m12.367-12.062a5.226 5.226 0 0 0-9.825-1.132"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.188 22.807a5.915 5.915 0 1 0 7.923-3.32m-17.798 3.262A52.274 52.274 0 0 0 4.5 20.554a45.853 45.853 0 0 1 8.113-.239m.786-.681a19.98 19.98 0 0 1-.726-8.307a22.297 22.297 0 0 1 7.993 3.058m2.021-1.035a11.71 11.71 0 0 1 1.4-7.415a13.767 13.767 0 0 0 1.111 5.279"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="3.643" d="m12.992 19.952l-1.503-1.328m10.001-4.8l-.43-1.683"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.307 20.195c.26-.232.747-.686 1.132-1.024c.446-.39 1.55-.15 1.55-.15l9.94-8.84a1.702 1.702 0 0 0 .141-2.404a1.627 1.627 0 0 0-2.301-.027h0l-9.94 8.842s.11 1.124-.332 1.52c-.472.423-1.351 1.203-1.351 1.203M41.023 9.7l-8.255 7.343"/>`),
		g.Group(children),
	)
}
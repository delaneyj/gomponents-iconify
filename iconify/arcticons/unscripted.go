package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unscripted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.728 34.418s-3.427-3.369-4.54-6.743m-1.05.758a8.507 8.507 0 0 0 1.88-5.55c-.172-3.206 1.475-7.655 1.475-7.655s-1.457-.806-3.05 2.897s-3.298 3.302-3.298 3.302"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.823 20.684a11.299 11.299 0 0 1-5.058-4.628m-.161 1.206l-.351-7.305c.212-.147 2.158-5.764-.317-6.457m0 0l-2.402 6.997s-.58 7.344-.438 7.739s.42 1.124.42 1.124"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.481 11.194s-1.475-.361-1.966 2.18c-1.013 5.24-.514 5.801.22 6.829"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.78 12.28s-1.84-1.02-2.515 2.954s-.324 2.97.385 4.215"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.966 17.167s-.803.07-1.305-1.788s-1.513-3.031-1.973-3.031s1.109 6.348 1.109 6.348s2.193 3.632 3.93 6.06m1.022 1.206c2.284 2.085 8.997 9.168 10.854 16.9m6.125-8.444s.69 1.517-1.836 5.027s-4.187 3.823-4.289 3.419m-6.289-31.643l2.318 1.047a2.16 2.16 0 0 1-2.236.661"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.776 14.024c-.05 3.224-1.52 8.565-3.753 9.291a9.032 9.032 0 0 0-5.311 6.047c-.894 3.003-3.526 9.703-3.526 9.703c3.55 5.29 6.808 5.435 6.808 5.435s.685-7.803 2.636-8.405s5.58-3.276 5.981-5.233s1.545-4.07 3.097-4.877s-.386-2.072-2.69-.532s-2.952 3.94-3.999 4.526"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.795 30.128a42.284 42.284 0 0 1 4.135-3.02m-13.032 3.974c1.095-.68 5.215-3.562 6.898-3.372m-6.092.36a12.102 12.102 0 0 0 1.23-2.833m5.139-.679a9.62 9.62 0 0 0 2.01-2.103m-3.566.66a6.252 6.252 0 0 0 2.405-2.502m-2.405 2.502a24.53 24.53 0 0 0 1.28-4.42"/>`),
		g.Group(children),
	)
}
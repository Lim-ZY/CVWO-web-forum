import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";

export const config = {
  matcher: ['/((?!_next/static|_next/image|favicon.ico).*)'],
}

export function middleware(request: NextRequest) {
  const cookieIsPresent: boolean = request.cookies.has(process.env.COOKIE_NAME || "Letters_JWT");
  const path: string = request.nextUrl.pathname;
  if (!cookieIsPresent && path !== '/') {
    return NextResponse.redirect(new URL('/', request.url));
  } else if (cookieIsPresent && path === '/') {
    return NextResponse.redirect(new URL('/t', request.url));
  } else {
    return NextResponse.next();
  }
}

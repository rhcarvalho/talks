class CustomAuthentication(authentication.BaseAuthentication):
    def authenticate(self, request):
        token = request.META.get("HTTP_X_TOKEN")
        if not token:
            return None
        try:
            user = get_user_by_token(token)
        except UserNotFound:
            raise exceptions.AuthenticationFailed
        return (user, "custom_auth")

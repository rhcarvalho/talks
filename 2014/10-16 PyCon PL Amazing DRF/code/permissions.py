class RandomAccess(BasePermission):
    def has_permission(self, request, view):
        return random.random() > 0.5


class IsOwner(BasePermission):
    def has_object_permission(self, request, view, obj):
        return obj.owner == request.user

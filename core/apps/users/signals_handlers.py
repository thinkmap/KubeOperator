from django.dispatch import receiver
from django_auth_ldap.backend import populate_user
from gunicorn.config import User

from users.models import Profile


@receiver(populate_user)
def on_ldap_create_user(sender, user, ldap_user, **kwargs):
    if user and user.username not in ['admin']:
        exists = Profile.objects.filter(user_id=user.id)
        if not exists:
            Profile.objects.create(user_id=user.id, source=Profile.USER_SOURCE_LDAP)

from locust import HttpUser, task


class TestC(HttpUser):

    payload = {
        "message": True,
    }

    @task
    def main(self):
        self.client.get('/')
        self.client.post('/', self.payload)

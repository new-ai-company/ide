from e2b import Sandbox


def main():
    explicit_api_key = Sandbox(id="Nodejs", api_key="YOUR_API_KEY")
    explicit_api_key.close()


main()

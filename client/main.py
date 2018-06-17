from jsonrpcclient.exceptions import ReceivedErrorResponse
from jsonrpcclient.http_client import HTTPClient
from models import Person


if __name__ == '__main__':
    client = HTTPClient('http://jsonrpc-server:5000/rpc')

    # this request will fail as 'surname' field can't be blank
    print("- TEST #1:\n")

    try:
        response = client.request(
            'DemoRPCService.CreatePerson',
            name='Gianfranco',
            surname='',
            age=36
        )
    except ReceivedErrorResponse as e:
        print("got exception %d: %s" % (e.code, e.message))

    # successfully create a person
    print("\n- TEST #2:\n")

    response = client.request(
        'DemoRPCService.CreatePerson',
        name='Gianfranco',
        surname='Reppucci',
        age=36
    )
    a_person = Person(response)

    print("person created with id: %s" % a_person.id)

    # get person from id
    print("\n- TEST #3:\n")

    response = client.request(
        'DemoRPCService.GetPerson',
        id=a_person.id
    )
    person_clone = Person(response)

    print("got person with id: %s, has name: %s and surname: %s" %
          (person_clone.id, person_clone.name, person_clone.surname))
